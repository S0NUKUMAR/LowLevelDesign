package main

import 
(
"fmt"
"time"
"example.com/bookmyshow/domain"
"sync"
"example.com/bookmyshow/repository"
"example.com/bookmyshow/service"
)

func buildScreen()*domain.Screen{
	var seats []*domain.Seat

	rows:=[]string{"A","B"}
	for _, row := range rows {
		for i := range 5 {
			t:= domain.SILVER
			if row == "B" {
				t = domain.GOLD
			}

			seats = append(seats, &domain.Seat{
				ID  : fmt.Sprintf("%s%d", row, i+1),
				Row : row,
				Number: i+1,
				Type:t,
			})
		}
	}
	return domain.NewScreen("screen-1","screen 1", seats)
}

func main(){
	/* User creation */
	alice:=domain.NewUser("user-alice", "Alice", "alice@example.com")
	bob:=domain.NewUser("user-bob", "Bob", "bob@example.com")


	fmt.Println("=============User Creation Done===========")
	fmt.Printf("%+v\n%+v\n\n", alice, bob)


	/*Theatre , Screen and Seat Creation*/
	theatre:= domain.NewTheatre("theatre-1", "PVR Forum","Pune")
	screen:=buildScreen()

	theatre.AddScreen(screen)

	fmt.Println("==============Theatre Created==============")

	fmt.Printf("%+v\n\n",theatre)

	
	movie:= domain.NewMovie("movie-1", "Intersteller", "English", "Sci-Fi", 90*time.Minute )
	start:= time.Now().Add(2*time.Hour)	
	end:= start.Add(movie.Duration)

	prices:= map[domain.SeatType]float64{
		domain.GOLD:350,
		domain.SILVER:200,
	}
	
	show:= domain.NewShow("show-1", movie, theatre, screen, start, end, prices)

	fmt.Println("==============Show Created==============")
	fmt.Printf("%+v\n\n", show)


	showRepo:= repository.NewInMemoryShowRepository()
	showRepo.Save(show)

	bookingRepo := repository.NewInMemoryBookingRepository()
	payment:= &service.MockPaymentGateway{}

	bookingSvc := service.NewBookingService(showRepo, bookingRepo, payment)

	fmt.Println("=== Scenario 1: two users race for the same seat ===")
	seatInContention := []string{"A1"}
	var wg sync.WaitGroup
	results := make(chan string, 2)

	race := func(user *domain.User) {
		defer wg.Done()
		if err := bookingSvc.HoldSeats(show.ID, user.ID, seatInContention); err != nil {
			results <- fmt.Sprintf("%s FAILED to hold A1: %v", user.Name, err)
			return
		}
		booking, err := bookingSvc.ConfirmBooking(show.ID, user.ID, seatInContention)
		if err != nil {
			results <- fmt.Sprintf("%s held A1 but failed to confirm: %v", user.Name, err)
			return
		}
		results <- fmt.Sprintf("%s BOOKED A1 successfully (booking %s, amount %.2f)", user.Name, booking.ID, booking.Amount)
	}

	wg.Add(2)
	go race(alice)
	go race(bob)
	wg.Wait()
	close(results)
	for r := range results {
		fmt.Println(r)
	}

	fmt.Println()
	fmt.Println("=== Scenario 2: hold expires if not confirmed in time ===")
	if err := bookingSvc.HoldSeats(show.ID, bob.ID, []string{"B3"}); err != nil {
		fmt.Println("hold failed:", err)
	} else {
		fmt.Println("Bob holds B3. Simulating time passing beyond the hold window...")
		state, _ := show.SeatState("B3", time.Now())
		fmt.Printf("Immediately after hold: status=%s heldBy=%s\n", state.Status, state.HeldBy)

		future := time.Now().Add(domain.DefaultHoldDuration + time.Second)
		state, _ = show.SeatState("B3", future)
		fmt.Printf("After hold window elapses: status=%s (seat is available again)\n", state.Status)

		// Alice can now grab it.
		if err := bookingSvc.HoldSeats(show.ID, alice.ID, []string{"B3"}); err != nil {
			fmt.Println("Alice failed to hold B3:", err)
		} else {
			fmt.Println("Alice successfully holds B3 after Bob's hold lapsed.")
		}
	}

	fmt.Println()
	fmt.Println("=== Scenario 3: remaining availability ===")
	fmt.Println("Available seats:", show.AvailableSeats(time.Now()))
		
}
