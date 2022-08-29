import 'dart:convert';
import 'package:analizador_lexicografico/table_elements.dart';
import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';
import 'dart:io';
import 'package:http/http.dart' as http;

void main() {
  runApp(const Analizador());
}

class Analizador extends StatelessWidget {
  const Analizador({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      title: "Analizador Lexicografico",
      home: HomeAnalizador(),
    );
  }
}

class HomeAnalizador extends StatefulWidget {
  const HomeAnalizador({Key? key}) : super(key: key);

  @override
  State<HomeAnalizador> createState() => _HomeAnalizadorState();
}

class _HomeAnalizadorState extends State<HomeAnalizador> {
  List<Tabla1Data> tabla1Cells = [];
  List<Tabla2Data> tabla2Cells = [];
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Analizador Lexicografico"),
      ),
      body: SingleChildScrollView(
        child: SizedBox(
          width: double.infinity,
          child: Wrap(
            alignment: WrapAlignment.center,
            spacing: 200,
            runSpacing: 50.0,
            children: <Widget>[
              Column(
                children: [
                  const Text(
                    "Tabla 1",
                    style: TextStyle(fontSize: 25),
                  ),
                  tablaLexico(tabla1Cells)
                ],
              ),
              Column(
                children: [
                  const Text(
                    "Tabla 2",
                    style: TextStyle(fontSize: 25),
                  ),
                  tablaTokens(tabla2Cells)
                ],
              ),
            ],
          ),
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () async {
          bool validRoute = await dialogOpen(context);
          if (validRoute) {
            circularProgress(context);
            List<Tabla1Data> tabla1 = [];
            List<Tabla2Data> tabla2 = [];
            Uri url1 = Uri.parse("http://localhost:8001/tabla1");
            var response = await http.get(url1);
            List<dynamic> decode = json.decode(response.body);
            for (dynamic tableElement in decode) {
              tabla1.add(Tabla1Data.fromJson(tableElement));
            }
            Uri url2 = Uri.parse("http://localhost:8001/tabla2");
            var response2 = await http.get(url2);
            List<dynamic> decode2 = json.decode(response2.body);
            for (dynamic tableElement in decode2) {
              tabla2.add(Tabla2Data.fromJson(tableElement));
            }
            Navigator.pop(context);
            setState(() {
              tabla1Cells = tabla1;
              tabla2Cells = tabla2;
            });
          }
        },
        backgroundColor: Colors.blue,
        child: const Icon(Icons.add),
      ),
    );
  }

  tablaLexico(List<Tabla1Data> tabla1) {
    return DataTable(
        columns: const <DataColumn>[
          DataColumn(label: Text("Nombre")),
          DataColumn(label: Text("Linea")),
          DataColumn(label: Text("Columna")),
          DataColumn(label: Text("Tipo 1")),
          DataColumn(label: Text("Tipo 2")),
          DataColumn(label: Text("Tipo 3"))
        ],
        rows: tabla1
            .map((item) => DataRow(cells: [
                  DataCell(Text(item.nombre)),
                  DataCell(Text(item.linea)),
                  DataCell(Text(item.columna)),
                  DataCell(Text(item.tipo1)),
                  DataCell(Text(item.tipo2)),
                  DataCell(Text(item.tipo3)),
                ]))
            .toList());
  }

  tablaTokens(List<Tabla2Data> tabla2) {
    return DataTable(
        columns: const <DataColumn>[
          DataColumn(label: Text("Token")),
          DataColumn(label: Text("#IdToken")),
          DataColumn(label: Text("Lexema generador")),
        ],
        rows: tabla2
            .map((item) => DataRow(cells: [
                  DataCell(Text(item.token)),
                  DataCell(Text(item.idToken.toString())),
                  DataCell(Text(item.lexema))
                ]))
            .toList());
  }
}

Future<bool> dialogOpen(BuildContext context) async {
  NavigatorState navState = Navigator.of(context);
  TextEditingController ruta = TextEditingController();
  final formKey = GlobalKey<FormState>();
  bool validRoute = false;

  await showDialog(
      context: context,
      builder: ((context) {
        return AlertDialog(
          title: const Text("Abrir archivo"),
          actions: [
            Form(
                key: formKey,
                child: TextFormField(
                  controller: ruta,
                  enabled: true,
                  decoration:
                      const InputDecoration(label: Text("Ruta del archivo")),
                  validator: ((value) {
                    if (value == null || value == "") {
                      return "Por favor ingresa una ruta";
                    } else {
                      return null;
                    }
                  }),
                )),
            Row(
              children: <Widget>[
                TextButton(
                    onPressed: () async {
                      FilePickerResult? result = await FilePicker.platform
                          .pickFiles(allowedExtensions: ['messi']);
                      if (result != null) {
                        File file = File(result.files.single.path!);
                        ruta.text = file.path;
                      } else {}
                    },
                    child: const Text("Abrir")),
                TextButton(
                    onPressed: () async {
                      if (formKey.currentState!.validate()) {
                        final isValidRuta = await File(ruta.text).exists();
                        if (isValidRuta) {
                          int statusCode = await sendFile(ruta.text);
                          if (statusCode == 200) {
                            validRoute = true;
                          }
                          validRoute = true;
                          navState.pop();
                        } else {
                          navState.pop();
                          archivoNoEncontrado(context);
                        }
                      }
                    },
                    child: const Text("Enviar"))
              ],
            )
          ],
        );
      }));
  return validRoute;
}

Future<dynamic> archivoNoEncontrado(
  BuildContext context,
) {
  return showDialog(
      context: context,
      barrierDismissible: true,
      builder: ((context) => const AlertDialog(
            title: Text("No se encontro el archivo"),
            actions: [Text("El archivo especificado no ha sido encontrado")],
          )));
}

Future<dynamic> circularProgress(BuildContext context) {
  return showDialog(
      context: context,
      barrierDismissible: true,
      builder: ((context) => const Center(
            child: CircularProgressIndicator(
              color: Colors.black,
            ),
          )));
}

Future<bool> validarArchivo(String value) async {
  return await File(value).exists();
}

Future<int> sendFile(String rute) async {
  Uri url = Uri.parse("http://localhost:8001/file");
  var request = http.MultipartRequest('POST', url);
  request.files.add(await http.MultipartFile.fromPath("file", rute));
  var response = await request.send();
  return response.statusCode;
}
