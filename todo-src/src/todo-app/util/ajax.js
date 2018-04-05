
export function GET(url) {
  return fetch(url).then(res => res.json())
}

export function POST(url,postParams) {
  const method = "POST"
  const headers = {
    'Content-Type': 'application/json'
  }
  const body = postParams
  //body: "{"data":"test"}" これおかしい
  console.log(url, {method, headers, body})
  return fetch(url, {method, headers, body}).then(res => res.json())
}