<?php

namespace App\Http\Controllers;

use Illuminate\Foundation\Bus\DispatchesJobs;
use Illuminate\Routing\Controller as BaseController;
use Illuminate\Foundation\Validation\ValidatesRequests;
use Illuminate\Foundation\Auth\Access\AuthorizesRequests;
use Illuminate\Foundation\Auth\Access\AuthorizesResources;

class Controller extends BaseController 
{
    protected function respError($err, $code = 500) {
        return response()->json(['error' => $err], $code);
    }

    protected function respSuccess($data, $code = 200) {
        return response()->json(['data' => $data], $code);
    }

}
