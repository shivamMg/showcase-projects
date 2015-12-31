$(document).ready(function(){
  var selectedTags = [];

  $(".button-collapse").sideNav();

  $(".chip").click(function(){
    var tag = $(this).text();
    var index = $.inArray(tag, selectedTags);
    // If `tag` in `selectedTags`
    if (index != -1){
      // Remove `tag` from `selectedTags`
      selectedTags.splice(index, 1);
    } else {
      // Add `tag` to `selectedTags`
      selectedTags.push(tag);
    }

    highlightTags(selectedTags);
  });

  function highlightTags(tagsList){
    $(".chip").each(function(){
      var ele = $(this);
      var tag = ele.text();
      if ($.inArray(tag, tagsList) != -1){
        ele.css("background", "#d32f2f");
      } else {
        ele.css("background", "#3e2723");
      }
    });
  }
});
