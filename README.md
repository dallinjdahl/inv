# inv
Command line stock portfolio manager

## Usage
    inv {-p portfolio-name | stock-ticker...}
## Portfolio format
The portfolio is stored in the home directory under the name `.(portfolio-name).port`.  It's a simple headerless CSV,
with the first field containing the stock ticker, and the second containing the number of stocks.  For example:
    FB, 4
    INTC, 2
    GOOG, 1
If this were stored in a the file `.dallin.port`, then I could run `inv -p dallin` and get the current value of my portfolio
and the amount it has increased (or decreased as the case may be) today.
