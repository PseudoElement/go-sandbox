const video = document.getElementById('video-tag');
const qualityInput = document.getElementById('quality')

const CHUNKS_SIZE = 86
let loadedChucnksNum = 0
let chunks = []
let timeoutId = null
let videoTiming = 0;

function checkValidQuality() {
  if (
    qualityInput.value !== '144' && 
    qualityInput.value !== '360' && 
    qualityInput.value !== '720' && 
    qualityInput.value !== '1080'
  ) {
    alert('Invalid quality!')
    throw new Error('Invalid quality!')
  }
}

function loadChunks() {
        chunks = [];
        checkValidQuality();
        const socket = new WebSocket(`ws://localhost:8080/api/v1/streaming/load-video?quality=${qualityInput.value}p&fileName=sea.mp4`)
        socket.binaryType = "blob"

        socket.onerror = function() {
            console.log('Close on error')
            socket.close()
        }
        socket.onclose = function() {
            console.log('Close on close')
            socket.close()
        }
        socket.onmessage = async function (msg) {
          chunks.push(msg.data)
          
          if (timeoutId) return;

          timeoutId = setTimeout(() => {
                const blob = new Blob(chunks, {type: 'mp4'})

                // console.log('Blob ==> ', msg.data)
                // console.log('Chunks ==> ', blob)
                console.log('video.currentTime ==> ', videoTiming)

                videoTiming = video.currentTime

                displayBlob(blob, video)

                video.currentTime = videoTiming
                video.play()
                loadedChucnksNum++
                timeoutId = null
          }, 2_000)

            // can pass simply msg.data without `type: "mp4"` assignment
        };
}

async function loadVideo() {
  fetch(
    "http://localhost:8080/api/v1/streaming/video", 
    {method: "GET", headers: {"Content-Type": "video/mp4"}}
  ).then((res) => res.blob())
    .then(blob => {
      displayBlob(blob, video)
      console.log('loadVideo_Blob ==> ', blob)
    })
    .catch(err => console.log('loadVideo_err ==> ', err)) 
}

/**
 * @param {Blob} blob
 * @param {HTMLVideoElement} videoEl 
 * @returns {void}
 */
function displayBlob( blob, videoEl ) {
  const newObjectUrl = URL.createObjectURL( blob );
  console.log('newObjectUrl ==> ', newObjectUrl);
  // const arr = [blob]
      
  // URLs created by `URL.createObjectURL` always use the `blob:` URI scheme: https://w3c.github.io/FileAPI/#dfn-createObjectURL
  const oldObjectUrl = videoEl.currentSrc;
  if( oldObjectUrl && oldObjectUrl.startsWith('blob:') ) {
      // It is very important to revoke the previous ObjectURL to prevent memory leaks. Un-set the `src` first.
      // See https://developer.mozilla.org/en-US/docs/Web/API/URL/createObjectURL
      videoEl.src = ''; // <-- Un-set the src property *before* revoking the object URL.
      URL.revokeObjectURL( oldObjectUrl );
  }

  // Then set the new URL:
  videoEl.src = newObjectUrl;

  // And load it:
  // videoEl.load(); // https://developer.mozilla.org/en-US/docs/Web/API/HTMLMediaElement/load
}

function fetchGreet() {
    fetch(
        "http://localhost:8080/api/v1/greeting",
        {method: "GET", headers: {"Content-Type": "application/json"}}
    ).then((res) => res.json())
     .then(data => console.log('Fetch_resp ==> ', data))
     .catch(err => console.log('Fetch_err ==> ', err)) 
}
