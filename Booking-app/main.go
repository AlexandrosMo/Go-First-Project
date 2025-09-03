// ΣΗΜΕΙΩΣΗ ΓΙΑ ΤΟΥΣ ΤΥΠΟΥΣ ΣΤΗ Go:
// Στη Go κάθε μεταβλητή πρέπει να έχει ΤΥΠΟ (π.χ. int, string).
// Αυτό χρησιμοποιείται για τρεις λόγους:
//
// 1. ΑΣΦΑΛΕΙΑ: Ο τύπος ορίζει τι τιμές επιτρέπονται. Αν βάλεις λάθος τιμή (π.χ. κείμενο σε int), το πρόγραμμα δεν τρέχει.
// 2. ΑΠΟΔΟΣΗ: Ο μεταγλωττιστής ξέρει ακριβώς πόσο χώρο μνήμης να κρατήσει και πώς να χειριστεί τη μεταβλητή.
// 3. ΚΑΘΑΡΟΤΗΤΑ: Ο τύπος δείχνει σε εμάς τι περιέχει η μεταβλητή (π.χ. string = κείμενο, int = αριθμός).
//
// Παράδειγμα:
// var userName string = "Tom"   // string → μπορεί να κρατήσει κείμενο
// var tickets int = 50          // int → μπορεί να κρατήσει ακέραιο αριθμό
//
// Αν βάλουμε λάθος τύπο, η Go μας σταματάει από το να τρέξουμε λάθος πρόγραμμα.

package main // Δηλώνει ότι αυτό το αρχείο ανήκει στο πακέτο "main". Σε εκτελέσιμα προγράμματα Go απαιτείται το πακέτο main.

import "fmt" // Φορτώνει το πακέτο "fmt" που παρέχει συναρτήσεις εκτύπωσης/εισόδου-εξόδου (π.χ. Printf, Println).

func main() { // Η συνάρτηση main είναι το σημείο εκκίνησης: από εδώ ξεκινά η εκτέλεση.

	var conferenceName string = "Go conference" // Δηλώνει μεταβλητή τύπου string (κείμενο). Τιμή: "Go conference".
	const conferenceTickets int = 50            // Δηλώνει σταθερά (const) τύπου int (ακέραιος). Δεν αλλάζει. Το σύνολο εισιτηρίων.
	var remainingTickets int = 50               // Μεταβλητή ακέραιου για τα διαθέσιμα εισιτήρια. Ξεκινά από 50.

	remainingTickets = -1 // Δίνουμε νέα τιμή στη μεταβλητή. Η Go το επιτρέπει γιατί είναι int, αλλά λογικά είναι λάθος (αρνητικά εισιτήρια).

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T \n",
		conferenceTickets, remainingTickets, conferenceName)
	// %T εκτυπώνει τον ΤΥΠΟ κάθε τιμής. Εδώ: int, int, string.

	fmt.Printf("welcome to %v booking application.\n", conferenceName)
	// %v εκτυπώνει την ΤΙΜΗ. Εδώ θα εμφανιστεί "Go conference".

	fmt.Printf("We have total of %v tickets and %v are still available.\n",
		conferenceTickets, remainingTickets)
	// Εκτυπώνει το σύνολο εισιτηρίων (50) και τα διαθέσιμα (-1).

	fmt.Println("Get your tickets here to attend")
	// Εκτυπώνει κείμενο και αλλάζει γραμμή αυτόματα.

	var userName string // Δηλώνουμε string (κείμενο) χωρίς τιμή. Προεπιλεγμένη τιμή = "" (κενό).
	var userTickets int // Δηλώνουμε int (ακέραιο) χωρίς τιμή. Προεπιλεγμένη τιμή = 0.

	// ask user for their name
	// Σχόλιο προγραμματιστή: δεν παίρνουμε πραγματικά είσοδο από τον χρήστη,
	// απλά βάζουμε χειροκίνητα τιμές ως παράδειγμα.

	userName = "Tom" // Αποθηκεύουμε "Tom" στη μεταβλητή userName.
	userTickets = 2  // Αποθηκεύουμε 2 στη μεταβλητή userTickets.

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
	// Εκτυπώνει: "User Tom booked 2 tickets."
}
