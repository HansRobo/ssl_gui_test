import { createApp } from 'vue';
import App from './App.vue';

const app = createApp(App);

// WebSocket接続を確立します。
const ws = new WebSocket('ws://localhost:8080');

// WebSocketのonmessageイベントハンドラを設定します。
ws.onmessage = (event) => {
  const shapes = JSON.parse(event.data);
  app.config.globalProperties.$shapes = shapes;
  app.component('App').methods.drawShapes();
};

app.mount('#app');
