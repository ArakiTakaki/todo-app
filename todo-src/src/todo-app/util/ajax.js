
export function GET(url) {
  return fetch(url).then(res => res.json())
}

export function POST(url,postParams) {
  const obj = postParams
  const method = "POST"
  let body
  for(let key in postParams){
    let value = encodeURIComponent(postParams[key])
    body += key + '=' + value + '&'
  }
  const headers = {
    'Accept': 'application/json',
    'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8'
  }

  console.log("POST TEST")
  console.log(body)
  console.log(url, {method, headers, body})
  console.log("POST TEST")

  return fetch(url, {method, headers, body}).then(res => res.json())
}