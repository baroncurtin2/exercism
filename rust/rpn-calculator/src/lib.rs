#[derive(Debug)]
pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    inputs.iter()
        .try_fold(vec![], |mut stack, input| {
            match *input {
                CalculatorInput::Value(x) => Some(x),
                _ => apply_operator(&mut stack, input)
            }
                .map(|x| {
                    stack.push(x);
                    stack
                })
        })
        .filter(|stack| stack.len() == 1)
        .map(|stack| stack[0])
}


fn apply_operator(values: &mut Vec<i32>, operator: &CalculatorInput) -> Option<i32> {
    values.pop()
        .and_then(|y| values.pop().map(|x| match operator {
            CalculatorInput::Add => x + y,
            CalculatorInput::Subtract => x - y,
            CalculatorInput::Multiply => x * y,
            CalculatorInput::Divide => x / y,
            _ => panic!("Invalid operator")
        }))
}
