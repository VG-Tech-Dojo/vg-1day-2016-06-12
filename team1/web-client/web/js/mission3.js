function changeBackground(colors) {
  var gradation =  'linear-gradient(45deg'
  for (var i = 0; i < colors.length; i++) {
    gradation += ', ' + colors[i];
  }
  gradation += ');'
  console.log(gradation);
  $('div.wrapper').css('background', gradation);
}
