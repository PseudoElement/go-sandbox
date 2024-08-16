async function wait(ms) {
    return new Promise((res) => setTimeout(res, ms));
}

const intervals = new Map([]);
let idCounter = 1;

function doInterval(callback, ms, endMs) {
    const thisId = idCounter;
    let startTime = Date.now();
    idCounter++;

    (async function interval() {
        intervals.set(thisId, this);
        while (intervals.has(thisId)) {
            const currentTime = Date.now();
            if (endMs && currentTime - startTime >= endMs) {
                removeInterval(thisId);
                return
            }
            await wait(ms);
            callback();
            if (!intervals.has(thisId)) return
        }
    })()

    return thisId; 
}

function removeInterval(id) {
    intervals.delete(id);
}

const doLog = () => console.log('Log!');
const intervalId = doInterval(doLog, 1000, 5000);

// setTimeout(() => removeInterval(intervalId), 5_001)