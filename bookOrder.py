
import json


files = (
    "genesis.json",
	"exodo.json",
	"levitico.json",
	"numeros.json",
	"deuteronomio.json",
	"josue.json",
	"jueces.json",
	"rut.json",
	"1_samuel.json",
	"2_samuel.json",
	"1_reyes.json",
	"2_reyes.json",
	"1_cronicas.json",
	"2_cronicas.json",
	"esdras.json",
	"nehemias.json",
	"ester.json",
	"job.json",
	"salmos.json",
	"proverbios.json",
	"eclesiastes.json",
	"cantares.json",
	"isaias.json",
	"jeremias.json",
	"lamentaciones.json",
	"ezequiel.json",
	"daniel.json",
	"joel.json",
	"oseas.json",
	"amos.json",
	"abdias.json",
	"jonas.json",
	"miqueas.json",
	"nahum.json",
	"habacuc.json",
	"sofonias.json",
	"hageo.json",
	"zacarias.json",
	"malaquias.json",
	"mateo.json",
	"marcos.json",
	"lucas.json",
	"juan.json",
	"hechos.json",
	"romanos.json",
	"1_corintios.json",
	"2_corintios.json",
	"galatas.json",
	"efesios.json",
	"filipenses.json",
	"colosenses.json",
	"1_tesalonicenses.json",
	"2_tesalonicenses.json",
	"1_timoteo.json",
	"2_timoteo.json",
	"tito.json",
	"filemon.json",
	"hebreos.json",
	"santiago.json",
	"1_pedro.json",
	"2_pedro.json",
	"1_juan.json",
	"2_juan.json",
	"3_juan.json",
	"judas.json",
	"apocalipsis.json",
)


order = 1
for file in files:
    if file.endswith(".json"):
        with open(f"data/json/{file}", "r", encoding="utf-8") as f:
            data = json.load(f)
        data["bookOrder"] = order
        order += 1
        with open(f"data/json2/{file}", "w", encoding="utf-8") as f:
            json.dump(data, f, ensure_ascii=False, indent=4)
            