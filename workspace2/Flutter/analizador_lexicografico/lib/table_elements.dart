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
      : nombre = json['nombre'],
        linea = json['linea'],
        columna = json['columna'],
        tipo1 = json['tipo1'],
        tipo2 = json['tipo2'],
        tipo3 = json['tipo3'];

  Map<String, dynamic> toJson() => {
        'nombre': nombre,
        'linea': linea,
        'columna': columna,
        'tipo1': tipo1,
        'tipo2': tipo2,
        'tipo3': tipo3
      };
}

class Tabla2Data {
  final String token;
  final String idToken;
  final String lexema;

  Tabla2Data(this.token, this.idToken, this.lexema);

  Tabla2Data.fromJson(Map<String, dynamic> json)
      : token = json['token'],
        idToken = json['idToken'],
        lexema = json['lexema'];

  Map<String, dynamic> toJson() =>
      {'token': token, 'idToken': idToken, 'lexema': lexema};
}
