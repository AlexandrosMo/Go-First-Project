// ΣΗΜΕΙΩΣΗ ΓΙΑ ΤΟΥΣ ΤΥΠΟΥΣ ΣΤΗ Go:
// Στη Go κάθε μεταβλητή έχει ΤΥΠΟ (π.χ. int, uint, string).
// - string = κείμενο (π.χ. "Tom", "hello")
// - int    = ακέραιος αριθμός (θετικός ή αρνητικός)
// - uint   = ακέραιος χωρίς αρνητικές τιμές (μόνο 0 και πάνω)
// Οι τύποι χρησιμοποιούνται για ασφάλεια, απόδοση και καθαρότητα κώδικα.

package main // Κάθε εκτελέσιμο πρόγραμμα Go ξεκινάει με το πακέτο "main".

import "fmt" // Πακέτο "fmt" = εργαλεία για εκτυπώσεις και ανάγνωση από το terminal.

func main() { // Η εκτέλεση ξεκινάει πάντα από τη συνάρτηση main.

	var conferenceName string = "Go conference" // Δηλώνουμε μεταβλητή τύπου string (κείμενο) με αρχική τιμή "Go conference".
	const conferenceTickets int = 50            // Δηλώνουμε σταθερά (const) τύπου int. Το σύνολο εισιτηρίων = 50 και δεν αλλάζει ποτέ.
	var remainingTickets uint = 50              // Δηλώνουμε μεταβλητή τύπου uint (ακέραιος ≥0). Τα διαθέσιμα εισιτήρια ξεκινούν από 50.

	fmt.Printf("conferenceTickets is %T,remainingTickets is %T,conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName)
	// Εκτυπώνουμε τους ΤΥΠΟΥΣ (όχι τις τιμές). Το %T δείχνει τον τύπο: int, uint, string.

	fmt.Printf("welcome to %v booking application.\n", conferenceName)
	// %v = εκτύπωσε την ΤΙΜΗ. Εδώ βάζει το "Go conference" μέσα στο κείμενο.

	fmt.Printf("We have total of %v tickets and %v  are still available.\n", conferenceTickets, remainingTickets)
	// Εκτυπώνουμε το σύνολο (50) και τα διαθέσιμα (50 αρχικά).

	fmt.Println("Get your tickets here to attend")
	// Println = απλή εκτύπωση με αλλαγή γραμμής.

	// Δηλώνουμε μεταβλητές για τα στοιχεία του χρήστη, χωρίς αρχικές τιμές.
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Ζητάμε είσοδο από τον χρήστη. Χρησιμοποιούμε fmt.Scan για να διαβάσουμε από το terminal.

	fmt.Println("Enter your name:") // Εμφανίζουμε μήνυμα.
	fmt.Scan(&firstName)            // &firstName = πέρασε τη "διεύθυνση μνήμης" ώστε η τιμή να αποθηκευτεί στη μεταβλητή.

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName) // Ο χρήστης γράφει το επίθετό του.

	fmt.Println("Enter your email:")
	fmt.Scan(&email) // Ο χρήστης γράφει το email του.

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets) // Ο χρήστης γράφει πόσα εισιτήρια θέλει.

	remainingTickets = remainingTickets - userTickets
	// Μειώνουμε τα διαθέσιμα εισιτήρια ανάλογα με την κράτηση.

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n",
		firstName, lastName, userTickets, email)
	// Εκτυπώνουμε προσωπικό μήνυμα επιβεβαίωσης.

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	// Ενημερώνουμε πόσα εισιτήρια απομένουν.
}
