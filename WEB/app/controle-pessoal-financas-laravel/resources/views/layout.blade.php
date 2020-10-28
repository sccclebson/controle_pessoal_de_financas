<!DOCTYPE html>
<html lang="pt-BR">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <link rel="stylesheet" href="css/bootstrap-reboot.min.css" />
    <link rel="stylesheet" href="css/bootstrap.min.css" />
    <link rel="stylesheet" href="css/{{ $css }}.css" />

    <title>Controle Pessoal de Finanças</title>
</head>

<body>
    <header>
        @yield('cabecalho')
    </header>

    <main>
        @yield('conteudo')
    </main>

    <footer>
        @yield('rodape')
    </footer>

    <script src="js/jquery-3.5.1.min.js"></script>
    <script src="js/bootstrap.min.js"></script>
    @yield('script')
</body>

</html>
