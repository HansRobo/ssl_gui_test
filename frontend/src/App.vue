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
      const canvas = document.getElementById('canvas');
      const ctx = canvas.getContext('2d');
      ctx.clearRect(0, 0, canvas.width, canvas.height);
      this.shapes.forEach(shape => {
        switch (shape.type) {
          case 'line':
            this.drawLine(ctx, shape);
            break;
          case 'arc':
            this.drawArc(ctx, shape);
            break;
        }
      });
    },
    drawLine(ctx, shape) {
      ctx.beginPath();
      ctx.moveTo(shape.x1, shape.y1);
      ctx.lineTo(shape.x2, shape.y2);
      ctx.lineWidth = shape.thickness;
      ctx.strokeStyle = 'black';
      ctx.stroke();
    },
    drawArc(ctx, shape) {
      ctx.beginPath();
      ctx.arc(shape.centerX, shape.centerY, shape.radius, shape.startAngle, shape.endAngle);
      ctx.lineWidth = shape.thickness;
      ctx.strokeStyle = 'black';
      ctx.stroke();
    }
  },
  mounted() {
    this.drawShapes();
  }
};
</script>

<style>
#canvas {
  border: 1px solid black;
}
</style>
