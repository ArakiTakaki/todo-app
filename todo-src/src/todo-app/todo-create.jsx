import React from 'react'


const error_log = (e) =>{
  console.log('todo-create:POST error')
  console.log(e)
}

// props
//    this.props.sheet_id
export default class TodoCreate extends React.Component{

  constructor(props){
    super(props)
    this.textParam = this.textParam.bind(this)
    this.post = this.post.bind(this)
  }

  post(url,postParams){
    const obj = postParams
    const method = "POST"
    //ここ解読してちゃんと理解する。
    const body = Object.keys(obj).map((key)=>key+"="+encodeURIComponent(obj[key])).join("&");
    const headers = {
      'Accept': 'application/json',
      'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8'
    }
    console.log(body)
    console.log(url, {method, headers, body})
    
    fetch(url, {method, headers, body})
      .then((res)=> res.json())
      .then(console.log)
      .catch(e => {
        console.log("POST")
      })
  }

  render(){ 
    return(
      <section>
        <button onClick={this.textParam}>[+]</button>
        <input type="text" name="todo_content" ref="todo" />
      </section>
      )
  }

  textParam(e){
    const url = '/api/create/todo'
    const post = {
      sheet : this.props.sheet_id,
      data  : this.refs.todo.value
    }
    this.post(url,post)
  }

}