type SubCallback = (count: number) => void

interface Subscriber {
    id: number;
    onUpdate: SubCallback
    unsubscibe: () => void
}

class Timer {
    private _value: number = 0

    private subs: Subscriber[] = []

    private _timerId: NodeJS.Timeout | null = null

    constructor(private readonly delayMs: number) {}

    public subscribe(onUpdate: SubCallback): () => void {
        if (!this.subs.length) this.run()

        const id = Date.now()
        const unsubscibe = () =>  {
            this.subs = this.subs.filter(sub => sub.id !== id)
            if(!this.subs.length) this.stop()
        }

        const newSub: Subscriber =  {
            id,
            onUpdate,
            unsubscibe
        }
        this.subs.push(newSub)

        return unsubscibe
    }

    private next(): void {
        this.subs.forEach(sub => sub.onUpdate(this._value))
        this._value++
    }

    private run(): void {
        this._timerId = setInterval(() => {
            this.next()
        }, this.delayMs)
    }

    private stop(): void {
        if (!this._timerId) return

        clearInterval(this._timerId)
        this._timerId = null
    }
}

export function timerRxJs(delayMs: number): Timer {
    return new Timer(delayMs)
}

const timer$ = timerRxJs(1_000)

const unsub1 = timer$.subscribe((idx) => console.log('First sub value: ', idx))
setTimeout(() => unsub1(), 10_000)

let unsub2: () => void
setTimeout(() =>  {
    unsub2 = timer$.subscribe((idx) => console.log('Seconds sub value: ', idx))
    setTimeout(() => unsub2(), 3_000)
}, 3_000)

setTimeout(() => {
    timer$.subscribe((idx) => console.log('After 14 secs value: ', idx))
}, 14_000)