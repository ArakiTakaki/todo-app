import React from 'react'
import View from './sheets-view.jsx'
import {GET} from './util/ajax'

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
    let json = GET("/api/todo/sheets")
    json.then( d => this.setState({data: d}) )
 }
}