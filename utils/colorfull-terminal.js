class StylishLog {
  msg = "";
  codes = "";

  // ANSI color codes
  static colors = {
    red: "\x1b[31m",
    green: "\x1b[32m",
    yellow: "\x1b[33m",
    blue: "\x1b[34m",
    purple: "\x1b[35m",
    cyan: "\x1b[36m",
    reset: "\x1b[0m",
  };

  constructor(msg) {
    this.msg = msg;
  }

  log() {
    console.log(`${this.codes}${this.msg}${StylishLog.colors.reset}`);
  }

  color(color) {
    this.codes += StylishLog.colors[color] || StylishLog.colors.purple;
    return this;
  }
}

new StylishLog("Hello").color("cyan").log();
new StylishLog("World").color("yellow").log();
