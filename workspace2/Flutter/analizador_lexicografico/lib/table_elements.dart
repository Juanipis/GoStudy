class Tabla1Data {
  final String nombre;
  final String linea;
  final String columna;
  final String tipo1;
  final String tipo2;
  final String tipo3;

  Tabla1Data(this.nombre, this.linea, this.columna, this.tipo1, this.tipo2,
      this.tipo3);

  Tabla1Data.fromJson(Map<String, dynamic> json)
      : nombre = json['name'],
        linea = json['line'],
        columna = json['numSimbFila'],
        tipo1 = json['t1'],
        tipo2 = json['t2'],
        tipo3 = json['t3'];

  Map<String, dynamic> toJson() => {
        'name': nombre,
        'line': linea,
        'numSimbFila': columna,
        't1': tipo1,
        't2': tipo2,
        't3': tipo3
      };
}

class Tabla2Data {
  final String token;
  final String idToken;
  final String lexema;

  Tabla2Data(this.token, this.idToken, this.lexema);

  Tabla2Data.fromJson(Map<String, dynamic> json)
      : token = json['Token'],
        idToken = json['IdToken'],
        lexema = json['LexemaGenerador'];

  Map<String, dynamic> toJson() =>
      {'Token': token, 'IdToken': idToken, 'LexemaGenerador': lexema};
}

class TablaAritmetica {
  final String expresion;
  final String linea;
  final String simboloInicio;
  final String simboloFinal;

  TablaAritmetica(
      this.expresion, this.linea, this.simboloInicio, this.simboloFinal);

  TablaAritmetica.fromJson(Map<String, dynamic> json)
      : expresion = json['exp'],
        linea = json['linea'],
        simboloInicio = json['simInicio'],
        simboloFinal = json['simFinal'];

  Map<String, dynamic> toJson() => {
        'exp': expresion,
        'linea': linea,
        'simInicio': simboloInicio,
        'simFinal': simboloFinal
      };
}
