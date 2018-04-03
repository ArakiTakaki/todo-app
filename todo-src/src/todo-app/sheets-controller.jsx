import React from 'react'
import View from './sheets-view.jsx'

// スタブ
//import list from './dammy/test-sheets.json'
// スタブ

// 情報群の取得処理
export default class SheetsController extends React.Component {
  constructor(props){
    super(props)
    this.state = {
      data: []
    }
    this.serchSheets = this.serchSheets.bind(this)
  }
  render(){
   
    const URI = CREATE_URI()
    // 結合処理様
    this.serchSheets(URI)

    return <View items={this.state.data}/>
  }

  serchSheets(uri){
    fetch(
      uri,
      { method: 'GET' }
    ).then( res=> {
      return res.json()
    }).then( d =>{
      this.setState({data: d})
    }).catch( e => {
      console.log("URI : " + uri)
      console.log("todo_listのajax処理に失敗しました")
      console.log(e)
    })

  }
}

const CREATE_URI = () =>{
  return "/api/todo/sheets"
}