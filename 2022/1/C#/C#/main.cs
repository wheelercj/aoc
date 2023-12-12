// https://adventofcode.com/2022/day/1

main();

static void main()
{
    part1();
    part2();
}

static void part1()
{
    //string content = File.ReadAllText("../../../test-input.txt");
    //int? expectedResult = 24000;

    string content = File.ReadAllText("../../../input.txt");
    int? expectedResult = 69836;

    List<int> allCalories = new() { 0 };
    foreach (string line in content.Split("\r\n"))
    {
        if (line.Length > 0)
            allCalories[^1] += int.Parse(line);
        else
            allCalories.Add(0);
    }

    allCalories.Sort();
    int mostCalories = allCalories.Last();

    Console.WriteLine($"Part 1 result: {mostCalories}");
    if (expectedResult != null && expectedResult != mostCalories)
    {
        throw new Exception($"Incorrect result: {mostCalories}");
    }
}

static void part2()
{
    //string content = File.ReadAllText("../../../test-input.txt");
    //int? expectedResult = 45000;

    string content = File.ReadAllText("../../../input.txt");
    int? expectedResult = 207968;

    List<int> allCalories = new() { 0 };
    foreach (string line in content.Split("\r\n"))
    {
        if (line.Length > 0)
            allCalories[^1] += int.Parse(line);
        else
            allCalories.Add(0);
    }

    allCalories.Sort();
    int top3Calories = allCalories.GetRange(allCalories.Count() - 3, 3).Sum();

    Console.WriteLine($"Part 2 result: {top3Calories}");
    if (expectedResult != null && expectedResult != top3Calories)
    {
        throw new Exception($"Incorrect result: {top3Calories}");
    }
}
