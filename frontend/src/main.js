import { createApp } from 'vue';
import App from './App.vue';

const app = createApp(App);

const ws = new WebSocket('ws://localhost:8080');
ws.onmessage = (event) => {
  const shapes = JSON.parse(event.data);
  app.config.globalProperties.$shapes = shapes;
};

app.mount('#app');
