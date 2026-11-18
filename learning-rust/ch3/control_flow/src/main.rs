fn main() {
    let number = 8;
    classic_if_else(number);

    looping();

    while_func();
    iterating_over_array([10, 20, 30, 40, 50]);

    let a = [5; 10];

    let mut sum = 0;

    for x in a {
        sum += x;
    }

    println!("{:?}", a);
    println!("{sum}");
}

fn classic_if_else(number: i32) {
    if number % 4 == 0 {
        println!("number is divisible by 4");
    } else if number % 3 == 0 {
        println!("number is divisible by 3");
    } else if number % 2 == 0 {
        println!("number is divisible by 2");
    } else {
        println!("number is not divisible by 4, 3, or 2");
    }
}

fn looping() {
    let mut counter = 0;

    let result = loop {
        counter += 1;

        if counter == 10 {
            break counter * 2;
        }
    };

    println!("The result is {result}");
}

fn while_func() {
    let mut number = 3;
    while number != 0 {
        println!("{number}!");

        number -= 1;
    }
    println!("LIFTOFF!!!");
}

fn iterating_over_array(input: [i32; 5]) {
    for element in input {
        println!("the value is: {element}");
    }
}
