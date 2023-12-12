import pytest


DAY = 0


def main():
    pass



with open(f"{DAY}/input.txt", "r") as file:
    INPUT = file.read()
main()


test_data = (
    ("", ""),
    ("", ""),
)


@pytest.mark.parametrize("input_,expected_output", test_data)
def test_main(input_, expected_output) -> None:
    assert main(input_) == expected_output
