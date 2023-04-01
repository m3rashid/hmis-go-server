### enum Dosage

How much in a day

| Name | Value |    Description    |
| ---- | :---: | :---------------: |
| OD   |   0   |    Once Daily     |
| BD   |   1   |    Twice Daily    |
| BDE  |   2   |    Twice Daily    |
| TD   |   3   | Three Times Daily |
| QD   |   4   | Four Times Daily  |
| FD   |   5   | Five Times Daily  |
| HD   |   6   |  Six Times Daily  |
| SD   |   7   | Seven Times Daily |

<br />

Times of day

| Name | Value | Description  |
| ---- | :---: | :----------: |
| BM   |   0   | Before meals |
| AM   |   1   | After meals  |

<br />

How much in a week

| Name | Value |    Description     |
| ---- | :---: | :----------------: |
| OW   |   0   |    Once Weekly     |
| BW   |   1   |    Twice Weekly    |
| TW   |   2   | Three Times Weekly |
| QW   |   3   | Four Times Weekly  |
| FW   |   4   | Five Times Weekly  |
| HW   |   5   |  Six Times Weekly  |
| SW   |   6   | Seven Times Weekly |

Dosage consists of a four digit number:

- First digit is the number of times in a day (if weekly, this is 1)
- Second digit is the time of the dose (before/ after meal)
- Third digit is the number of days of dose in a week (usually 7 if the medicine is to be taken daily)
- Fourth digit is the total duration of the medicine (in days)

### enum Consumables MedicineType

| Name | Value |        Description         |
| ---- | :---: | :------------------------: |
| TAB  |   0   |           Tablet           |
| SYR  |   1   |           Syrup            |
| INJ  |   2   |         Injection          |
| SAL  |   3   |           Saline           |
| SYR  |   4   |          Syringe           |
| DRS  |   5   |          Dressing          |
| GLO  |   6   |           Gloves           |
| TAP  |   7   |            Tape            |
| PIP  |   8   |            Pipe            |
| TST  |   9   |        Testing Kits        |
| BND  |  10   |          Bandages          |
| ALW  |  11   |       Alcohol Wipes        |
| GZP  |  12   |         Gauze Pads         |
| MTC  |  13   |        Medical Tape        |
| CTN  |  14   |        Cotton Balls        |
| STG  |  15   |       Sterile Gloves       |
| SYR  |  16   |          Syringes          |
| NDL  |  17   |          Needles           |
| IVC  |  18   |        IV Catheters        |
| CTH  |  19   |       Catheter Bags        |
| WND  |  20   |      Wound Dressings       |
| URB  |  21   |         Urine Bags         |
| OXM  |  22   |        Oxygen Masks        |
| NCS  |  23   |       Nasal Cannulas       |
| SUC  |  24   |     Suction Catheters      |
| ETT  |  25   |     Endotracheal Tubes     |
| TTT  |  26   |     Tracheostomy Tubes     |
| FDT  |  27   |       Feeding Tubes        |
| CTL  |  28   |     Central Line Kits      |
| DBS  |  29   |   Disposable Bed Sheets    |
| SBL  |  30   |      Surgical Blades       |
| SCL  |  31   |          Scalpels          |
| SPS  |  32   |      Surgical Sponges      |
| EKG  |  33   |         EKG Paper          |
| SPC  |  34   |    Specimen Containers     |
| SHC  |  35   |     Sharps Containers      |
| SDP  |  36   |      Surgical Drapes       |
| STK  |  37   |        Suture Kits         |
| CST  |  38   |     Casting Materials      |
| XRF  |  39   |        X-Ray Films         |
| CCM  |  40   | CT Scan Contrast Materials |
| MCM  |  41   |   MRI Contrast Materials   |
| UGM  |  42   |      Ultrasound Gels       |
| DBB  |  43   |        Dental Bibs         |
| FCM  |  44   |         Face Masks         |
| FCS  |  45   |        Face Shields        |
| SWP  |  46   |      Sanitizing Wipes      |
| ITV  |  47   |       IV Tubing Sets       |
| IVS  |  48   |      IV Solution Bags      |

### enum NonConsumables MedicineType

| Name | Value |                Description                |
| ---- | :---: | :---------------------------------------: |
| HBD  |   0   |               Hospital Bed                |
| WCH  |   1   |                Wheelchair                 |
| CRU  |   2   |                 Crutches                  |
| CAN  |   3   |                   Canes                   |
| WKL  |   4   |                  Walkers                  |
| MBS  |   5   |             Mobility Scooters             |
| BPR  |   6   |          Blood Pressure Monitors          |
| NEB  |   7   |                Nebulizers                 |
| OXY  |   8   |           Oxygen Concentrators            |
| ECG  |   9   |    Electrocardiography (ECG) Machines     |
| AED  |  10   |              Defibrillators               |
| STB  |  11   |              Surgical tables              |
| OPL  |  12   |             Operating lights              |
| ANM  |  13   |            Anesthesia machines            |
| PTL  |  14   |               Patient lifts               |
| RAD  |  15   |            Radiology equipment            |
| USM  |  16   |            Ultrasound machines            |
| MRI  |  17   | Magnetic resonance imaging (MRI) machines |
| CTM  |  18   |     Computed tomography (CT) scanners     |
| XRM  |  19   |              X-ray machines               |
| ENS  |  20   |                Endoscopes                 |
| ALC  |  21   |                Autoclaves                 |
| STE  |  22   |          Sterilization equipment          |
| PNM  |  23   |             Patient monitors              |
| EEG  |  24   |   Electroencephalography (EEG) machines   |
| EMG  |  25   |      Electromyography (EMG) machines      |
| HAI  |  26   |               Hearing aids                |
| EYG  |  27   |                Eye glasses                |
| CTL  |  28   |              Contact lenses               |
| DCH  |  29   |               Dental chairs               |
| DRL  |  30   |               Dental drills               |
| DSC  |  31   |              Dental scalers               |
| DXM  |  32   |           Dental x-ray machines           |
| OIM  |  33   |            Orthopedic implants            |
| PTO  |  34   |         Prosthetics and orthotics         |
