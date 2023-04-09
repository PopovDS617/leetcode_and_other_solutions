interface NodeItem {
  name: string;
  children: NodeItem[];
}

const tree: NodeItem = {
  name: 'root',
  children: [
    {
      name: 'node_1',
      children: [
        { name: 'node_1.1', children: [] },
        { name: 'node_1.2', children: [{ name: 'node_1.2.1', children: [] }] },
      ],
    },
    { name: 'node_2', children: [{ name: 'node_2.1', children: [] }] },
    { name: 'node_3', children: [] },
  ],
};

function getLeafItemNames(item: NodeItem) {
  let node = item;

  for (let curr in item) {
    node = curr;

    if (typeof node === 'array') {
      for (let i = 0; i < node.length; i++) {}
    }
  }

  throw new Error('Not implemented');
}

// console.log(getLeafItemNames(tree)); // ['node_1.1', 'node_1.2.1', 'node_2.1', 'node_3']
