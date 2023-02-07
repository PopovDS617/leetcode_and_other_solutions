//find a lonely integer in array of duplicates

let duplicates:any =[]
  duplicates=a.filter((num,index)=>{
    return a.indexOf(num)!==index
  })
  
  const result:number[]|number=a.filter(item=>{
    if(!duplicates.includes(item)){return item}
      
  })
  const final:number=result[0]
  
   return final
}
