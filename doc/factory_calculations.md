https://satisfactory.wiki.gg/wiki/Heavy_Modular_Frame

# calculation
given a target amount of output for a single factory, would like to calculate a **builder count** and **clockrate** to meet that output, where the clock rate is always below 1.0

procedure:

- take target value and divide by output per minute of 1 factory
- this gives a float number of factories to create the desired input if all factories are at 100% (N factories)
- round the number up, then do `<target value> / N` to get the output amount each factory would need to produce if had N factories
- since N was rounded up, this needed output amount per factory is guaranteed to be below the 100% rate of the factory
- do `<original output amount of factory> / <calculated scaled down amount per factory>` to get the clock rate for each N factory