// https://adventofcode.com/2022/day/3
main();

static void main()
{
    part1();
    part2();
}

static void part1()
{
    //string content = File.ReadAllText("../../../test-input.txt");
    //int? expected = 157;

    string content = File.ReadAllText("../../../input.txt");
    int? expected = 7727;

    int sum = 0;
    string[] lines = content.Split("\r\n", StringSplitOptions.RemoveEmptyEntries);
    foreach (string line in lines)
    {
        int mid = line.Length / 2;
        string left = line[..mid];
        string right = line[mid..];

        foreach (char ch in left)
        {
            if (right.Contains(ch))
            {
                sum += getPriority(ch);
                break;
            }
        }
    }

    Console.WriteLine(sum);
    if (expected != null && expected != sum)
        throw new Exception($"Incorrect sum: {sum}");
}

static void part2()
{
    //string content = File.ReadAllText("../../../test-input.txt");
    //int? expected = 70;

    string content = File.ReadAllText("../../../input.txt");
    int? expected = 2609;

    int sum = 0;
    string[] lines = content.Split("\r\n", StringSplitOptions.RemoveEmptyEntries);
    for (int i = 0; i < lines.Length - 2; i += 3)
    {
        string line1 = lines[i];
        string line2 = lines[i + 1];
        string line3 = lines[i + 2];

        char badge = ' ';
        foreach (char ch in line1)
        {
            if (line2.Contains(ch) && line3.Contains(ch))
            {
                badge = ch;
                break;
            }
        }
        
        sum += getPriority(badge);
    }

    Console.WriteLine(sum);
    if (expected != null && expected != sum)
        throw new Exception($"Incorrect sum: {sum}");
}

static int getPriority(char ch)
{
    if (ch >= 'a' && ch <= 'z')
        return ch - 'a' + 1;
    if (ch >= 'A' && ch <= 'Z')
        return ch - 'A' + 27;
    throw new Exception($"Invalid character: {ch}");
}
