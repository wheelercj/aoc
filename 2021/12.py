def main():
    graph = []
    for line in INPUT.splitlines():
        first, second = line.split('-')
        insert_edge(first, second, graph)

    for node in graph:
        if node[0] == 'start':
            start = node
            break

    paths = []
    walk_graph(start, [], paths)

    for path in paths:
        print(','.join(path))
    print(f'\n{len(paths) = }')  # 3738


def insert_edge(node1_name: str, node2_name: str, graph: list[list]) -> None:
    node1 = insert_node(node1_name, graph)
    node2 = insert_node(node2_name, graph)
    node1.append(node2)
    node2.append(node1)


def insert_node(node_name: str, graph: list[list]) -> list:
    for node in graph:
        if node[0] == node_name:
            return node
    new_node = [node_name]
    graph.append(new_node)
    return new_node


def walk_graph(node: list,
               this_path: list[str],
               all_paths: list[list[str]]) -> None:
    _this_path = list(this_path)  # copy
    _this_path.append(node[0])
    if node[0] == 'end':
        all_paths.append(_this_path)
        return
    for next_node in node[1:]:
        if next_node[0].isupper() or next_node[0] not in _this_path:
            walk_graph(next_node, _this_path, all_paths)

# The inputs here are not present in the public repo as requested
# https://adventofcode.com/2021/about

SAMPLE1 = '''\
'''

SAMPLE1_PATH_COUNT = 0

SAMPLE2 = '''\
'''

SAMPLE2_PATH_COUNT = 0

SAMPLE3 = '''\
'''

SAMPLE3_PATH_COUNT = 0

INPUT = '''\
'''

if __name__ == '__main__':
    main()
