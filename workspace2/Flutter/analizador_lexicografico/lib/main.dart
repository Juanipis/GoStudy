import 'dart:convert';
import 'dart:ffi';

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
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Analizador Lexicografico"),
      ),
      body: SizedBox(
        width: double.infinity,
        child: Wrap(
          alignment: WrapAlignment.center,
          spacing: 200,
          runSpacing: 4.0,
          children: <Widget>[
            Column(
              children: [const Text("Tabla 1"), tablaLexico(tabla1Cells)],
            ),
            Column(
              children: [const Text("Tabla 2"), tablaTokens()],
            ),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () async {
          bool validRoute = await dialogOpen(context);
          if (validRoute) {
            circularProgress(context);
            List<Tabla1Data> tabal1DataAct = await generateTabla1List();
            Navigator.pop(context);
            setState(() {
              tabla1Cells = tabal1DataAct;
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

  tablaTokens() {
    return DataTable(columns: const <DataColumn>[
      DataColumn(label: Text("Token")),
      DataColumn(label: Text("#IdToken")),
      DataColumn(label: Text("Lexema generador")),
    ], rows: const <DataRow>[
      DataRow(cells: <DataCell>[
        DataCell(Text("Hola")),
        DataCell(Text("Mundo")),
        DataCell(Text("Mundo"))
      ])
    ]);
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
          title: const Text("Abir archivo"),
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
                    child: const Text("Abir")),
                TextButton(
                    onPressed: () async {
                      if (formKey.currentState!.validate()) {
                        final isValidRuta = await File(ruta.text).exists();
                        if (isValidRuta) {
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

Future<List<Tabla1Data>> generateTabla1List() async {
  List<Tabla1Data> tabla1 = [];
  var url = Uri.https('api.npoint.io', '/e305395fbcf20817bfa2');
  var response = await http.get(url);
  Map<String, dynamic> decode = json.decode(response.body);
  //print(decode["tabla1"].runtimeType);
  for (dynamic tableElement in decode["tabla1"]) {
    tabla1.add(Tabla1Data.fromJson(tableElement));
  }
  return tabla1;
}
