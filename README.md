# quickwrite.xyz ![Stars](https://img.shields.io/github/stars/Simpson-Computer-Technologies-Research/quickwrite.xyz?color=brightgreen) ![Watchers](https://img.shields.io/github/watchers/Simpson-Computer-Technologies-Research/quickwrite.xyz?label=Watchers)
![Untitled](https://user-images.githubusercontent.com/75189508/187591268-184a9eac-9c93-4ebf-89c8-adaa9626b29a.png)

# About
This project was made using Svelte and Golang. QuickWrite iterates over the provided input and finds synonyms for every other word. QuickWrite has a caching system for both frontend and backend to ensure faster synonym queries. 

# How to use
**1.** Paste in or type in any text you want into the white box. 
<br>

**2.** Click refresh.
<br>

**3.** Hover over the purple text to display a dropdown containing up to seven synonyms.
<br>

**4.** Click on one of the synonyms and the word will be replaced

# API
## Usage

```
$ curl -X GET http://localhost:8000/synonyms?q=fast
```

```json
"Example Response": {
    ["agile","brisk","nimble","quick","rapid","swift"]
}
```

# License
MIT License

Copyright (c) 2022 Tristan Simpson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
