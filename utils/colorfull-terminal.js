class StylishLog {
    msg = ""
    styles = ""

    constructor(msg) {
        this.msg = msg
    }

    log() {
        console.log(`%c${this.msg}`, this.styles)
        this.msg = null;
        this.styles = null;
    }

    color(color) {
        this.styles += `color: ${color || 'pink'}; `
        return this
    }

    fontSize(fs) {
        this.styles += `font-size: ${fs || '20px'}; `
        return this
    }

    textDecoration(td) {
        this.styles += `text-decoration: ${td || 'none'}; `
        return this
    }
}

new StylishLog('Hello').color('purple').fontSize('40px').textDecoration('underline').log()
new StylishLog('World').color('green').fontSize('40px').textDecoration('dotted').log()
