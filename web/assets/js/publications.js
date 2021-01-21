$('#new-publication').on('submit', createPublication)
$(document).on('click', '.like-publication', likePublication)
$(document).on('click', '.deslike-publication', deslikePublication)
$('#update-publication').on('click', updatePublication)
$('.delete-publication').on('click', deletePublication)

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
    Swal.fire(
      'Sucesso!',
      'Publicação criada com sucesso!',
      'success'
    ).then(function(){
      window.location = "/feed"
    })
  }).fail(function(err){
    console.log(err)
    Swal.fire(
      'Ops...',
      'Erro ao criar a publicação!',
      'error'
    )
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
    Swal.fire(
      'Ops...',
      'Erro ao curtir a publicação!',
      'error'
    )
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
    Swal.fire(
      'Ops...',
      'Erro ao descurtir a publicação!',
      'error'
    )
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
    Swal.fire(
      'Sucesso!',
      'Publicação atualizada com sucesso!',
      'success'
    ).then(function(){
      window.location = "/feed"
    })
  }).fail(function(){
    Swal.fire(
      'Ops...',
      'Erro ao atualizar a publicação!',
      'error'
    )
  }).always(function(){
    $('#update-publication').prop('disabled', false)
  })
}

function deletePublication(e){
  e.preventDefault();

  Swal.fire({
    title: 'Atenção!',
    text: 'Deseja excluir essa publicação?',
    showCancelButton: true,
    cancelButtonText: 'Cancelar',
    icon: 'warning'
  }).then(function(confirm){
    if(!confirm.value) return

    const elm = $(e.target)
    const publication = elm.closest('div')
    const publicationId = publication.data('publication-id')

    elm.prop('disabled', true)

    $.ajax({
      url: `/publications/${publicationId}`,
      method: 'DELETE',
    }).done(function() {
      publication.fadeOut("slow", function(){
        $(this).remove()
      })
    }).fail(function(){
      Swal.fire(
        'Ops...',
        'Erro ao deletar a publicação!',
        'error'
      )
    })
  })

}