import requests

print(requests.get("http://127.0.0.1:8090", headers={
    "TargetUrl": "https://httpbin.org/deflate",
    "JA3": "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0",
    "UA": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0"
}, params={
    "hey": "yo",
    "hey[yo]": "{}@#Q)$*(_)&*(@#)&*(#ADHJKLAHJKL#:JKLRW:JKLFA GOWNO) \""
}, json={
    "hey": "yooo"
}).text)