# CSVexplorer


  ____ ______     __              _                     
 / ___/ ___\ \   / /____  ___ __ | | ___  _ __ ___ _ __ 
| |   \___ \\ \ / / _ \ \/ / '_ \| |/ _ \| '__/ _ \ '__|
| |___ ___) |\ V /  __/>  <| |_) | | (_) | | |  __/ |   
 \____|____/  \_/ \___/_/\_\ .__/|_|\___/|_|  \___|_|   
  by: guicapelleto         |_|

Versão: 1.5

*Menu de ajuda*

./csvexplorer [opções]

Opções:
--arquivo      -a   =C:\pasta\/01\arquivo.csv          Atribui um arquivo para leitura. Obs: usar \/ como espaço.
--limite       -l   =4                                 Limita o número dos resultados encontrados.
--ler-coluna   -lc  =1,6-8                             Limita a exibição dos resultados encontrados para as colunas mencionadas.
--ler-linha    -ll  =1,6-8                             Limita a exibição dos resultados encontrados para as linhas mencionadas.
--filtrar      -f   =palavra1,palavra2                 Realiza um filtro de texto, onde somente será exibido caso o texto mencionado for encontrado na linha.
--filtrarColuna-fc  =7:palavra1,palavra2               Realiza um filtro de texto, onde somente será exibido caso o texto mencionado for encontrado na linha.
--separador    -s   =;                                 Altera o padrão de vírgula como separação de campos por outro, no exemplo ponto e vírgula.
--break        -b                                      Quebra a exibição das colunas em linhas, ideal para ver separadamente os elementos de uma linha específica.
--verbose      -v                                      Ativa informações adicionais.
--contagem     -c                                      Exibe o total de linhas obtidas.
--unicos       -u                                      Não exibe linhas repetidas.
--silenciar    -S                                      Não exibe o resultado das linhas.
--help         -h                                      Exibe esse menu de ajuda.

Explorador de arquivos CSV.
