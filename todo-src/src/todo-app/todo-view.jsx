import React from 'react'

export default class TodoView extends React.Component{
  render(){
    
    var items = this.props.items
    var list = []
    for (var item of items ){
      list.push(
        <li key={item.id}>
          {item.checkd}
          {item.todo_name}
        </li>
      )
    }

    return (
      <ul>
        {list}
      </ul>
    )

  }
}