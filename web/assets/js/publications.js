$('#new-publication').on('submit', createPublication)
$('.like-publication').on('click', likePublication)

function createPublication(e){
  e.preventDefault();

  $.ajax({
    url: "/publications",
    method: 'POST',
    data: {
      title: $('#title').val(),
      content: $('#content').val(),
    }
  }).done(function() {
    window.location = "/feed"
  }).fail(function(){
    alert("Erro ao criar a publicação")
  })
}

function likePublication(e){
  e.preventDefault();

  const elm = $(e.target);
  const publicationId = elm.closest('div').data('publication-id')

  $.ajax({
    url: `/publications/${publicationId}/like`,
    method: "POST"
  }).done(function () {
    alert("Publicação Curtida")
  }).fail(function() {
    alert("Erro ao curtir")
  })

}

