const SECONDS_IN_A_DAY: u64 = 24 * 60 * 60;
fn main() {
    let x = 5;
    println!("The value of x is: {x}");

    println!("Number of seconds in a day : {SECONDS_IN_A_DAY}");

    // Shadowing
    {
        let x = 6;
        println!("The value of x is: {x}");
    }
    println!("The value of x is: {x}");

    // Tuples
    let tup = (500, 6.4, 1);

    let (x, y, z) = tup;

    println!("The value of x is: {x}");
    println!("The value of y is: {y}");
    println!("The value of z is: {z}");

    let t = ([1; 2], [3; 4]);

    let (a, _) = t;

    println!("{}", a[0] + t.1[0]);
}
