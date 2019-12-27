<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>To Do List</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link href='https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css' rel="stylesheet" type="text/css">
    <link href='https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.0.3/css/font-awesome.css' rel="stylesheet" type="text/css">
    <link href="{{ asset('assets/css/index.css') }}" rel="stylesheet" type="text/css">
    <script src='https://cdnjs.cloudflare.com/ajax/libs/jquery/3.2.1/jquery.min.js'></script>
    <script src='https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js'></script>
    <script src="{{ asset('assets/js/index.js') }}"></script>
    
</head>
<body>
<div class="page-content page-container" id="page-content">
    <div class="padding">
        <div class="row justify-content-center">
            <div class="col-lg-12">
                <div class="card px-3">
                    <div class="card-body">
                        <div class="add-items d-flex"> 
                            <input id="todo-input" type="text" class="form-control todo-list-input" placeholder="What do you need to do today?"> 
                            <button id="add-todo" class="add btn btn-primary font-weight-bold todo-list-add-btn">Add</button> 
                        </div>
                        <div class="list-wrapper">
                            <ul class="d-flex flex-column todo-list" id="todo-list">
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
    
</body>
</html>