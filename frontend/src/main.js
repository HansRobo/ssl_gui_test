import { createApp } from 'vue';
import App from './App.vue';

const app = createApp(App);

// WebSocket接続を確立します。
const ws = new WebSocket('ws://localhost:8080');

// WebSocketのonmessageイベントハンドラを設定します。
ws.onmessage = (event) => {
  const data = event.data.split(',');
  const type = data[0];
  let shape = null;

  if (type === 'line') {
    shape = {
      type: 'line',
      x1: parseFloat(data[1]),
      y1: parseFloat(data[2]),
      x2: parseFloat(data[3]),
      y2: parseFloat(data[4]),
      color: 'black'
    };
  } else if (type === 'circle') {
    shape = {
      type: 'circle',
      x: parseFloat(data[1]),
      y: parseFloat(data[2]),
      radius: parseFloat(data[3]),
      color: 'black'
    };
  }

  if (shape) {
    if (!app.config.globalProperties.$shapes) {
      app.config.globalProperties.$shapes = [];
    }
    app.config.globalProperties.$shapes.push(shape);
  }
};

app.mount('#app');
