DAY = 7


class File:
    def __init__(self, name, size):
        self.name = name
        self.size = size


class Folder:
    def __init__(self, path):
        self.path = path


def get_folder_path(folder_name: str, filesystem: dict[str, list[File|Folder]]) -> str:
    if folder_name == "/":
        return "/"
    for folder_path in filesystem:
        if folder_path.endswith("/" + folder_name):
            return folder_path
    raise ValueError


def get_folder_size(folder_path: str, filesystem: dict[str, list[File|Folder]]) -> int:
    size = 0
    for item in filesystem[folder_path]:
        if isinstance(item, File):
            size += item.size
        elif isinstance(item, Folder):
            size += get_folder_size(item.path, filesystem)
    return size


def parse_input(input_text: str) -> dict[str, list[File|Folder]]:
    current_dir = "/"
    filesystem = {"/": []}
    ignore_lines = 0
    lines = input_text.splitlines()[1:]
    for i, line in enumerate(lines):
        if ignore_lines:
            ignore_lines -= 1
            continue
        dollar_sign, command, *args = line.split()
        if command == "cd":
            folder_name = args[0]
            if folder_name == "..":
                current_dir = "/".join(current_dir.split("/")[:-1])
            else:
                if current_dir == "/":
                    current_dir = "/" + folder_name
                else:
                    current_dir = current_dir + "/" + folder_name
                if folder_name not in filesystem:
                    filesystem[current_dir] = []
        elif command == "ls":
            if filesystem[current_dir]:
                continue
            j = i + 1
            line = ""
            while not line.startswith("$") and j < len(lines):
                line = lines[j]
                if not line.startswith("$"):
                    ignore_lines += 1
                    first, second = line.split()
                    if first == "dir":
                        if current_dir == "/":
                            new_dir_path = "/" + second
                        else:
                            new_dir_path = current_dir + "/" + second
                        filesystem[new_dir_path] = []
                        filesystem[current_dir].append(Folder(new_dir_path))
                    else:
                        filesystem[current_dir].append(File(second, int(first)))
                    j += 1
    return filesystem


def find_best_folder_size(min_size: int, filesystem: dict[str, list[File|Folder]]) -> int:
    """Finds the smallest folder >= min_size and returns its size."""
    best_size = get_folder_size("/", filesystem)
    for path in filesystem:
        if path == "/":
            continue
        size = get_folder_size(path, filesystem)
        if size >= min_size and size < best_size:
            best_size = size
    return best_size


filesystem = parse_input(INPUT)
result = get_folder_size(get_folder_path("e", filesystem), filesystem)
assert result == 584, result
result = get_folder_size(get_folder_path("a", filesystem), filesystem)
assert result == 94853, result
result = get_folder_size(get_folder_path("d", filesystem), filesystem)
assert result == 24933642, result
result = get_folder_size(get_folder_path("/", filesystem), filesystem)
assert result == 48381165, result

used_space = get_folder_size("/", filesystem)
assert used_space == 48381165, used_space

TOTAL_FREE_SPACE = 70_000_000
NEEDED_FREE_SPACE = 30_000_000

free_space = TOTAL_FREE_SPACE - used_space
assert free_space == 21618835, free_space
needed_free_space = NEEDED_FREE_SPACE - free_space
assert needed_free_space == 8381165, needed_free_space
best_folder_size = find_best_folder_size(needed_free_space, filesystem)
assert best_folder_size == 24933642, best_folder_size

with open(f"{DAY}/input.txt", "r") as file:
    INPUT = file.read()
filesystem = parse_input(INPUT)
used_space = get_folder_size("/", filesystem)
free_space = TOTAL_FREE_SPACE - used_space
needed_free_space = NEEDED_FREE_SPACE - free_space
best_folder_size = find_best_folder_size(needed_free_space, filesystem)
assert best_folder_size != used_space  # probably
print(f"{best_folder_size = }")  # 545729
