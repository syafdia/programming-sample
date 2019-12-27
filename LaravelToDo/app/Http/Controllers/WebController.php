<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\Task;

class WebController extends Controller 
{
    public function index() 
    {
        return view('index');
    }

}