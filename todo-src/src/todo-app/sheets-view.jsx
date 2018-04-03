import React from 'react'
import {Link} from 'react-router-dom'

export default class Sheets extends React.Component {
  render(){
    var data = this.props.items
    var list = []
    for (var value of data){      
      list.push(
        <li key={value.id}>
          <Link to={"/todo/"+value.id}>
            {value.sheet_title}<span>{value.update_at}</span>
          </Link>
        </li>
      )
    }

    return(
        <ul>
          {list}
        </ul>
    )
  }
}
