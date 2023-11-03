const c = document.getElementById("console");
const urlParams = new URLSearchParams(window.location.search);
const password = urlParams.get("p");

let tid = 0;
let ws;

const t = document.getElementById("commandinput");
t.addEventListener("input", () => {
  const value = t.value;
  if (!ws) {
    return;
  }
  const command = value.split(" ").shift();
  ws.send(
    JSON.stringify({
      type: "find-command",
      data: {
        id: tid,
        command,
      },
    })
  );
  tid++;
});

const r = document.getElementById("commandrun");
r.addEventListener("click", () => {
  const command = t.value;
  ws.send(
    JSON.stringify({
      type: "command",
      data: command,
    })
  );
  t.value = "";
});

if (!password) {
  log(JSON.stringify({ type: "error", message: "Please input a password" }));
}

log(JSON.stringify({ type: "error", message: "Connecting..." }));
connect();

function connect() {
  ws = new WebSocket("ws://" + window.location.host);
  ws.onopen = function () {
    ws.send(
      JSON.stringify({
        type: "auth",
        data: password,
      })
    );
  };
  ws.onmessage = function (ev) {
    const msg = JSON.parse(ev.data);
    switch (msg.type) {
      case "sync": {
        clear();
        syncLog(msg.data.log);
        break;
      }
      case "log": {
        log(JSON.parse(msg.data));
        break;
      }
      case "playeradd":
        const player = msg.data;
      case "response":
        if (msg.data.data) {
          t.style["color"] = "white";
        } else {
          t.style["color"] = "red";
        }
    }
  };
  ws.onclose = function () {
    window.location.replace("/login");
  };
}

function syncLog(text) {
  const msgs = text.split("\n");
  msgs.forEach(function (m) {
    const msg = JSON.parse(m);
    log(msg);
  });
}

function log(msg) {
  if (msg.type == "chat") {
    msg = JSON.parse(msg.message);
    const msgs = [msg];
    if (msg.extra) {
      for (const m of msg.extra) {
        msgs.push(m);
      }
    }
    let txt = "";
    for (const m of msgs) {
      txt += `<a class="consoletext" style="color: ${
        m.color || "white"
      }">${m.text.replaceAll("<", "&lt;")}</a>`;
    }
    c.innerHTML += txt + "<br/>";
  } else {
    let type = "";
    switch (msg.type) {
      case "info": {
        type += `<a class="consoletext" style="color: #1F74E2">INFO</a>`;
        break;
      }
      case "debug": {
        type += `<a class="consoletext" style="color: cyan">DEBUG</a>`;
        break;
      }
      case "error": {
        type += `<a class="consoletext" style="color: red">ERROR</a>`;
        break;
      }
      case "warn": {
        type += `<a class="consoletext" style="color: yellow">WARN</a>`;
        break;
      }
    }
    let m = msg.message;
    if (m) {
      m = m.replaceAll("\n", "<br/>");
    }
    c.innerHTML += `<a class="consoletext" style="color: gray">${msg.time}</a> ${type}<a class="consoletext" style="color: white">: ${m}</a><br/>`;
  }
}

function clear() {
  c.innerHTML = "";
}
