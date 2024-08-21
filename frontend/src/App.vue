<template>
  <div id="app">
    <canvas id="canvas" width="800" height="600"></canvas>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      shapes: []
    };
  },
  methods: {
    drawShapes() {
      const canvas = this.$refs.canvas;
      const ctx = canvas.getContext('2d');
      ctx.clearRect(0, 0, canvas.width, canvas.height);
      this.shapes.forEach(shape => {
        switch (shape.type) {
          case 'circle':
            this.drawCircle(ctx, shape);
            break;
          case 'line':
            this.drawLine(ctx, shape);
            break;
          case 'rectangle':
            this.drawRectangle(ctx, shape);
            break;
          case 'text':
            this.drawText(ctx, shape);
            break;
        }
      });
    },
    drawCircle(ctx, shape) {
      ctx.beginPath();
      ctx.arc(shape.x, shape.y, shape.radius, 0, 2 * Math.PI);
      ctx.fillStyle = shape.color;
      ctx.fill();
    },
    drawLine(ctx, shape) {
      ctx.beginPath();
      ctx.moveTo(shape.x1, shape.y1);
      ctx.lineTo(shape.x2, shape.y2);
      ctx.strokeStyle = shape.color;
      ctx.stroke();
    },
    drawRectangle(ctx, shape) {
      ctx.fillStyle = shape.color;
      ctx.fillRect(shape.x, shape.y, shape.width, shape.height);
    },
    drawText(ctx, shape) {
      ctx.font = shape.font;
      ctx.fillStyle = shape.color;
      ctx.fillText(shape.text, shape.x, shape.y);
    }
  },
  mounted() {
    const canvas = this.$refs.canvas;
    const ctx = canvas.getContext('2d');
    this.drawShapes();
  },
  watch: {
    shapes: 'drawShapes'
  }
};
</script>

<style>
#canvas {
  border: 1px solid black;
}
</style>
