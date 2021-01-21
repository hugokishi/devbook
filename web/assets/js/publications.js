$('#new-publication').on('submit', createPublication)
$(document).on('click', '.like-publication', likePublication)
$(document).on('click', '.deslike-publication', deslikePublication)
$('#update-publication').on('click', updatePublication)

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

  elm.prop('disabled', true)

  $.ajax({
    url: `/publications/${publicationId}/like`,
    method: "POST"
  }).done(function () {
    const like_counter = elm.next('span')
    const qtd_likes = parseInt(like_counter.text())
    like_counter.text(qtd_likes + 1);

    elm.addClass('deslike-publication')
    elm.addClass('text-danger')
    elm.removeClass('like-publication')

  }).fail(function() {
    alert("Erro ao curtir")
  }).always(function() {
    elm.prop('disabled', false)
  })
}

function deslikePublication(e){
  e.preventDefault();

  const elm = $(e.target);
  const publicationId = elm.closest('div').data('publication-id')

  elm.prop('disabled', true)

  $.ajax({
    url: `/publications/${publicationId}/deslike`,
    method: "POST"
  }).done(function () {
    const like_counter = elm.next('span')
    const qtd_likes = parseInt(like_counter.text())
    like_counter.text(qtd_likes - 1);

    elm.removeClass('deslike-publication')
    elm.removeClass('text-danger')
    elm.addClass('like-publication')

  }).fail(function() {
    alert("Erro ao descurtir")
  }).always(function() {
    elm.prop('disabled', false)
  })
}

function updatePublication(e){
  $(this).prop('disabled', true)

  const publicationId = $(this).data('publication-id')
  
  $.ajax({
    url: `/publications/${publicationId}`,
    method: 'PUT',
    data: {
      title: $('#title').val(),
      content: $('#content').val(),
    }
  }).done(function() {
    alert("Atualizado")
  }).fail(function(){
    alert("Erro")
  }).always(function(){
    $('#update-publication').prop('disabled', false)
  })
}

