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
  }
  render(){
    return <View items={this.state.data}/>
  }
  // 初期化ajaxはここに書くらしい
  componentDidMount(){

    //DEBUG
    //this.setState({data: list})
    
    fetch(
      CREATE_URI(),
      { method: 'GET' }
    ).then( res=> {
      return res.json()
    }).then( d =>{
      this.setState({data: d})
    }).catch( e => {
      console.log("todo_listのajax処理に失敗しました")
      console.log(e)
    })
  }
}

const CREATE_URI = () =>{
  return "/api/todo/sheets"
}