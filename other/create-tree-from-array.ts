// from multiple paths

const toTree = (paths: string[]) => {
  const tree = {};
  for (const path of paths) {
    if (path) {
      let node = tree;
      const parts = path.split('/');

      for (const part of parts) {
        if (part) {
          //   if (node[part] === undefined || node[part] === null) {
          //     node[part] = {};
          //   }
          // }

          // node = node[part];
          node = node[part] ?? (node[part] = {});
        }
      }
    }
    return tree;
  }
};

const paths = ['src/App/index.js', 'public/icon.ico'];

// from single path

const dig = (path: string) => {
  let tree = {};
  const parts = path.split('/');

  let node = tree;
  for (let part of parts) {
    if (part) {
      if (node[part] === undefined || node[part] === null) {
        if (part === 'treasure') {
          node[part] = '10$';
        } else {
          node[part] = {};
        }
      }

      node = node[part];
    }
  }

  return tree;
};

const test = 'deep/deeper/deepest/treasure';
