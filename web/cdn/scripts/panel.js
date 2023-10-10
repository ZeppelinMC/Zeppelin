const c = document.getElementById("console");
const urlParams = new URLSearchParams(window.location.search);
const password = urlParams.get("p");
if (!password) {
  log(JSON.stringify({ type: "error", message: "Please input a password" }));
}

log(JSON.stringify({ type: "error", message: "Connecting..." }));
connect();

function connect() {
  const ws = new WebSocket("ws://" + window.location.host);
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
    console.log(msg.type);
    switch (msg.type) {
      case "sync": {
        clear();
        sync(msg.data);
        break;
      }
      case "log": {
        log(JSON.parse(msg.data));
        break;
      }
    }
  };
  ws.onclose = function () {
    window.location.replace("/login");
  };
}

function sync(text) {
  const msgs = text.split("\n");
  msgs.forEach(function (m) {
    const msg = JSON.parse(m);
    log(msg);
  });
}

function log(msg) {
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
      type += `<a class="consoletext" style="color: yellow">ERROR</a>`;
      break;
    }
  }
  c.innerHTML += `<a class="consoletext" style="color: gray">${msg.time}</a> ${type}<a class="consoletext" style="color: white">: ${msg.message}</a><br/>`;
}

function clear() {
  c.innerHTML = "";
}
