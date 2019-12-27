$(document).ready(function() {
  const STATUS_ACTIVE = 1;
  const STATUS_COMPLETED = 2;

  const $todoLi = $('#todo-list');
  const $todoInput = $('#todo-input');
  const $addTodoBtn = $('#add-todo');

  function ajax(httpMethod, route, data) {
      return new Promise((resolve, reject) => {
          const req = {
            type: httpMethod,
            url: route,
            contentType: 'application/json',
          };

          if (data) {
            req.data = JSON.stringify(data);
          }

          $.ajax(req)
            .done((resp) => resolve(resp))
            .fail((err) => reject(err));
      })
  }

  function fetchAllTasks() {
    return ajax('GET', '/tasks');
  }

  function updateTaskStatus(id, taskStatus) {
    return ajax('PUT', `/tasks/${id}`, { status: taskStatus});
  }

  function createTask(name) {
    return ajax('POST', `/tasks`, { name, status: STATUS_ACTIVE });
  }

  function deleteTask(id) {
    return ajax('DELETE', `/tasks/${id}`);
  }

  function renderTasksList() {
    fetchAllTasks().then((resp) => {
      $todoInput.val('');
      $todoLi.children().remove();
      $todoLi.append(resp.data.map((task) => renderTaskItem(task)));
    });
  }

  function renderTaskItem(task) {
    const $taskList = $(`
      <li class="${task.status === STATUS_COMPLETED ? 'completed' : ''}">
        <div class="form-check"> 
          <label class="form-check-label"> 
            <input class="checkbox" type="checkbox" ${task.status === STATUS_COMPLETED ? 'checked' : ''}>
            ${task.name}<i class="input-helper" />
          </label>
        </div>
        <i class="remove mdi mdi-close-circle-outline"/>
      </li>`);

    $taskList.find('input')
      .on('click', function() {
        updateTaskStatus(task.id, $(this).is(':checked') ? STATUS_COMPLETED : STATUS_ACTIVE)
          .then(() => renderTasksList());
      });

    $taskList.find('.remove')
      .on('click', (e) => {
        deleteTask(task.id)
        .then(() => renderTasksList());
      });

    return $taskList;
  }

  (function init() {
    $addTodoBtn.on('click',(e) => {
      e.preventDefault();
      createTask($todoInput.val().trim())
        .then(() => renderTasksList())
    });

    renderTasksList();
  })();

})