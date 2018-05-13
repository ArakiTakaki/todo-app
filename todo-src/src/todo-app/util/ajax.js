
export function GET(url) {
  return fetch(url).then(res => res.json())
}

export function POST(url,postParams) {
  const method = "POST"
  const obj = postParams
  const body = Object.keys(obj).map((key)=>key+"="+encodeURIComponent(obj[key])).join("&");
  const headers = {
    'Accept': 'application/json',
    'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8'
  }
  console.log(url, {method, headers, body})
  return fetch(url, {method, headers, body}).then(res => res.json())
}