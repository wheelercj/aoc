// https://adventofcode.com/2022/day/2

main();

static void main()
{
    part1();
    part2();
}

static void part1()
{
    //string content = File.ReadAllText("../../../test-input.txt");
    //int? expected = 15;

    string content = File.ReadAllText("../../../input.txt");
    int? expected = 13009;

    var theirChoices = new Dictionary<string, ChoicePoints>()
    {
        { "A", ChoicePoints.rock },
        { "B", ChoicePoints.paper },
        { "C", ChoicePoints.scissors },
    };
    var yourChoices = new Dictionary<string, ChoicePoints>()
    {
        { "X", ChoicePoints.rock },
        { "Y", ChoicePoints.paper },
        { "Z", ChoicePoints.scissors },
    };
    
    int totalScore = 0;
    string[] lines = content.Split("\r\n", StringSplitOptions.RemoveEmptyEntries);
    foreach (string line in lines)
    {
        string[] choices = line.Split(' ');
        ChoicePoints theirs = theirChoices[choices[0]];
        ChoicePoints yours = yourChoices[choices[1]];
        totalScore += getScorePart1(theirs, yours);
    }

    Console.WriteLine(totalScore);
    if (expected != null && expected != totalScore)
        throw new Exception($"Invalid totalScore: {totalScore}");
}

static void part2()
{
    //string content = File.ReadAllText("../../../test-input.txt");
    //int? expected = 12;

    string content = File.ReadAllText("../../../input.txt");
    int? expected = 10398;

    var theirChoices = new Dictionary<string, ChoicePoints>()
    {
        { "A", ChoicePoints.rock },
        { "B", ChoicePoints.paper },
        { "C", ChoicePoints.scissors },
    };
    var yourChoices = new Dictionary<string, ResultChoice>()
    {
        { "X", ResultChoice.lose },
        { "Y", ResultChoice.draw },
        { "Z", ResultChoice.win },
    };

    int totalScore = 0;
    string[] lines = content.Split("\r\n", StringSplitOptions.RemoveEmptyEntries);
    foreach (string line in lines)
    {
        string[] choices = line.Split(' ');
        ChoicePoints theirs = theirChoices[choices[0]];
        ResultChoice yours = yourChoices[choices[1]];
        totalScore += getScorePart2(theirs, yours);
    }

    Console.WriteLine(totalScore);
    if (expected != null && expected != totalScore)
        throw new Exception($"Invalid totalScore: {totalScore}");
}

static int getScorePart1(ChoicePoints theirs, ChoicePoints yours)
{
    if (yours == theirs)
        return (int)ResultPoints.draw + (int)yours;
    
    bool youWin = (
        yours == ChoicePoints.rock && theirs == ChoicePoints.scissors
        || yours == ChoicePoints.paper && theirs == ChoicePoints.rock
        || yours == ChoicePoints.scissors && theirs == ChoicePoints.paper
    );
    if (youWin)
        return (int)ResultPoints.win + (int)yours;

    bool youLose = (
        yours == ChoicePoints.rock && theirs == ChoicePoints.paper
        || yours == ChoicePoints.paper && theirs == ChoicePoints.scissors
        || yours == ChoicePoints.scissors && theirs == ChoicePoints.rock
    );
    if (youLose)
        return (int)ResultPoints.lose + (int)yours;
    throw new Exception("Invalid choice combination or choice handling.");
}

static int getScorePart2(ChoicePoints theirs, ResultChoice yours)
{
    if (yours == ResultChoice.draw)
        return (int)ResultPoints.draw + (int)theirs;
    var winPoints = new Dictionary<ChoicePoints, int>()
    {
        // { their choice, your points }
        { ChoicePoints.rock, (int)ChoicePoints.paper + (int)ResultPoints.win },
        { ChoicePoints.paper, (int)ChoicePoints.scissors + (int)ResultPoints.win },
        { ChoicePoints.scissors, (int)ChoicePoints.rock + (int)ResultPoints.win },
    };
    var losePoints = new Dictionary<ChoicePoints, int>()
    {
        // { their choice, your points }
        { ChoicePoints.rock, (int)ChoicePoints.scissors + (int)ResultPoints.lose },
        { ChoicePoints.paper, (int)ChoicePoints.rock + (int)ResultPoints.lose },
        { ChoicePoints.scissors, (int)ChoicePoints.paper + (int)ResultPoints.lose },
    };

    if (yours == ResultChoice.win)
        return winPoints[theirs];
    if (yours == ResultChoice.lose)
        return losePoints[theirs];
    throw new Exception("Invalid choice combination or choice handling.");
}

enum ResultPoints
{
    win = 6,
    lose = 0,
    draw = 3,
}

enum ChoicePoints
{
    rock = 1,
    paper,
    scissors,
}

enum ResultChoice
{
    lose,
    draw,
    win,
}
