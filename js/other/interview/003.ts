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

function getLeafItemNames(item ) {
  const result=[]
  
  const stack=[]
  
  
  stack.push(tree)
  
while(stack.length>0){
  
  const currentObj=stack.pop()
  
  if(currentObj.name){result.unshift(currentObj.name)}
  if(Array.isArray(currentObj.children)){
    
    currentObj.children.forEach(item=>{
      stack.push(item)
    })
    
  }
  else if(currentObj.children){stack.push(currentObj.children)}
  
  
  
}
   
  return result
}
 

// console.log(getLeafItemNames(tree)); // ['node_1.1', 'node_1.2.1', 'node_2.1', 'node_3']
