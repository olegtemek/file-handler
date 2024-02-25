export const UseFile = ()=>{


  const getTags = async () => {
    const req = await fetch(`${import.meta.env.VITE_API_URL}/v1/file/tags`, { method: 'GET' })
    const data = await req.json()
    return data.tags
  }
  const getFilesByTag = async (tag) => {
    const req = await fetch(`${import.meta.env.VITE_API_URL}/v1/file?tag=${tag}`, { method: 'GET' })
    const data = await req.json()

    return data.files
  }


  const getText = async (filepath) =>{
    const req =  await fetch(`${import.meta.env.VITE_API_URL}/${filepath}`, {method:"GET"})
    
    const data = await req.text()

    if (data == ""){
      return "Nothing...."
    }

    return data
    
  }
  const remove = async (id) =>{
    const req =  await fetch(`${import.meta.env.VITE_API_URL}/v1/file/${id}`, {method:"DELETE"})
    
    const data = await req.json()

    if(data.status == 200){
      return true
    }
    return false
    
  }


  return {
    getTags,
    getFilesByTag,
    getText,
    remove
  }


}